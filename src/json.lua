-- json 

local json = {}
local dkjson = require("BSE-DCS-Export/dkjson") or nil
local Logger = require("BSE-DCS-Export/logger")

function json:dump(obj)
    if dkjson == nil then        
        Logger:warning("Cannot find a suitable library to dump json, returning object as-is...")        
        return obj    
    end

    return dkjson.encode(obj)
end


return json