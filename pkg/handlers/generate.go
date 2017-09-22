package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	txtTemplate "text/template"

	"github.com/k8s-community/codegen/pkg/generator"
	"github.com/k8s-community/codegen/pkg/router"
	codeGenTemplate "github.com/k8s-community/codegen/pkg/template"
	"github.com/k8s-community/codegen/pkg/utils"
)

// TemplateData is data for service template
type TemplateData struct {
	ServiceName        string
	EnvPrefix          string
	ServiceDescription string
	ProjectPath        string
	ProjectURL         string
	HomePageURL        string
	OwnerName          string
	OwnerEmail         string

	RegistryURL    string
	Namespace      string
	Infrastructure string

	ChartEnvs map[string]ChartEnvironment
}

// ChartEnvironment defines chart params for special environment
type ChartEnvironment struct {
	RegistryHost          string
	ServiceHost           string
	CommonServicesHost    string
	Namespace             string
	TLSSecretName         string
	TLSServicesSecretName string
}

// GenerateCode handles requests for code generation
func (h *Handler) GenerateCode(c router.Control) {
	r := c.Request()

	templatePath := "templates/generate-success.html"

	htmlData := make(map[string]string)

	archivePath, err := h.generateServiceArchive(r)
	if err != nil {
		templatePath = "templates/generate.html"

		htmlData["err"] = fmt.Sprintf("%s", err)
		for key := range r.Form {
			htmlData[key] = r.Form.Get(key)
		}
	}

	t, err := template.ParseFiles("templates/layout.html", templatePath)
	if err != nil {
		h.logger.Fatalf("Cannot parse `generate.html` or `layout.html`: %s", err)
	}

	htmlData["link"] = strings.TrimLeft(archivePath, "/tmp")
	t.ExecuteTemplate(c, "layout", htmlData)
}

func (h *Handler) generateServiceArchive(r *http.Request) (string, error) {
	config, err := h.getGeneratorConfigFromRequest(r)
	if err != nil {
		h.logger.Infof("cannot get config from request: %s", err)
		return "", fmt.Errorf("There is a some problem with sent parameters.")
	}

	err = config.Validate()
	if err != nil {
		h.logger.Infof("invalid sent parameters: %s", err)
		return "", fmt.Errorf("Invalid sent parameters. Please check name of your service, service path.")
	}

	err = generator.GenerateCode(*config)
	if err != nil {
		h.logger.Infof("cannot generate code: %s with config %#v", err, *config)
		return "", fmt.Errorf("Cannot generate service: %s", err)
	}

	destFile := "/tmp/archive/" + fmt.Sprintf("%s.tar.gz", utils.RandomString(16))
	err = utils.CreateTarGzArchive(config.DestPath, destFile)
	if err != nil {
		h.logger.Infof("cannot generate code: %s with config %#v", err, *config)
		return "", fmt.Errorf("Cannot generate service: %s", err)
	}

	return destFile, nil
}

func (h *Handler) getGeneratorConfigFromRequest(r *http.Request) (*generator.Config, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	serviceName := r.Form.Get("service_name")

	data := &TemplateData{
		ServiceName: serviceName,
		EnvPrefix:   h.createEnvPrefixFromServiceName(serviceName),

		ServiceDescription: r.Form.Get("service_description"),
		ProjectPath:        r.Form.Get("project_path"),
		ProjectURL:         r.Form.Get("project_url"),
		HomePageURL:        r.Form.Get("homepage_url"),
		OwnerName:          r.Form.Get("owner_name"),
		OwnerEmail:         r.Form.Get("owner_email"),

		RegistryURL:    r.Form.Get("registry_url"),
		Namespace:      r.Form.Get("namespace"),
		Infrastructure: r.Form.Get("infrastructure"),
	}

	// TODO: get it from request
	devParams := ChartEnvironment{
		RegistryHost:          "registry.k8s.community",
		ServiceHost:           "codegen-dev.k8s.community",
		CommonServicesHost:    "services-dev.k8s.community",
		Namespace:             "dev",
		TLSSecretName:         "codegen-tls-secret",
		TLSServicesSecretName: "tls-secret",
	}

	// TODO: get it from request
	stableParams := ChartEnvironment{
		RegistryHost:          "registry.k8s.community",
		ServiceHost:           "codegen.k8s.community",
		CommonServicesHost:    "services.k8s.community",
		Namespace:             "stable",
		TLSSecretName:         "codegen-tls-secret",
		TLSServicesSecretName: "tls-secret",
	}

	data.ChartEnvs = make(map[string]ChartEnvironment)
	data.ChartEnvs["dev"] = devParams
	data.ChartEnvs["stable"] = stableParams

	srcDir := "code-templates/go-rest"
	destDir := fmt.Sprintf("/tmp/%s", utils.RandomString(16))

	replacePaths := make(map[string]string)
	replacePaths["cmd/k8sapp.go"] = "cmd/{[( .ServiceName )]}.go"

	return &generator.Config{
		AppName: r.Form.Get("service_name"),
		Config: codeGenTemplate.Config{
			SrcPath:    destDir,
			Data:       data,
			LeftDelim:  "{[(",
			RightDelim: ")]}",
			FuncMap: txtTemplate.FuncMap{
				"ToUpper": strings.ToUpper,
				"ToLower": strings.ToLower,
			},
		},
		SrcPath:      srcDir,
		DestPath:     destDir,
		ReplacePaths: replacePaths,
	}, nil
}

func (h *Handler) createEnvPrefixFromServiceName(serviceName string) string {
	prefix := strings.ToUpper(serviceName)

	// Env prefix must not consist of -
	prefix = strings.Replace(prefix, "-", "", -1)

	return prefix
}
