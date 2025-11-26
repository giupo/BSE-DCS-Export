-- Logger Wrapper, the original sucks big time.
local log = log or {}

local Logger = {}

function Logger:info(msg)
    log.write("BSE", log.INFO, msg)
end

function Logger:warning(msg)
    log.write("BSE", log.WARNING, msg)
end

function Logger:error(msg)
    log.write("BSE", log.ERROR, msg)
end

function Logger:debug(msg)
    log.write("BSE", log.DEBUG, msg)
end

return Logger