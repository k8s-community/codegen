function toggler(divId) {
    div = $("#"+divId)
    if (div.hasClass("hidden")) {
        div.removeClass("hidden")
    } else {
        div.addClass("hidden")
    }

    return false;
}