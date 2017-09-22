package generator

import "testing"

var appNameValidationFixtures = []struct {
	appName  string
	isPassed bool
}{
	{"myapp", true},
	{"my_app", false},
	{"3rdpts", false},
	{"my-app", true},
	{"my3app", true},
	{"MyApp", false},
	{"myApp", false},
	{"my.app", false},
	{"my:app", false},
	{"MY_APP", false},
	{"my/app", false},
	{"abcdefghabcdefghabcdefghabcdefg", true},
	{"abcdefghabcdefghabcdefghabcdefgha", false},
}

func TestAppNameValidationHandler(t *testing.T) {
	for _, expData := range appNameValidationFixtures {
		envConfig := &Config{
			AppName: expData.appName,
		}

		err := envConfig.validateServiceName()
		isPassed := (err != nil)

		if expData.isPassed == isPassed {
			t.Errorf("Incorrect validation of appName '%s', expected = %t, isError : %t", expData.appName, expData.isPassed, isPassed)
		}
	}
}
