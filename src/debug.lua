local DEBUG = true

local Logger = require("BSE-DCS-Export/logger")

local debug = {
    active = false
}

-- debug functions:

function debug:writeFile(path, text)
    local file = io.open(path, "w")
    if not file then
        Logger:error("Cannot open file " .. path)
        return false
    end

    file:write(text)
    file:close()
    return true
end


return debug