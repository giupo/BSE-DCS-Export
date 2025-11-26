local Logger = require("BSE-DCS-Export/logger")
local json = require("BSE-DCS-Export/json")

local udp = {
    ip =  "127.0.0.1",
    port = 6666
}

function udp:init(ip, port)
    local socket = require("socket")
    self.socket = socket.udp()
    self.socket:settimeout(0)
    self.socket:setoption("broadcast", true)
end

function udp:send(data)
    if self.socket == nil then
        Logger:warning("Cannot send, UDP not initizlized")
        return
    end

    local payload = (data and json:dump(data) or "")    
    local rc, err_msg = self.socket:sendto(payload, self.ip, self.port)

    if err_msg ~= nil then
        Logger:warning("Cannot send this motherfucker: " .. err_msg)
    end
end

function udp:close()
    if self.socket ~= nil then
        self.socket:close()
    end
end


return udp