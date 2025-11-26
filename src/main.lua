-- just to make the linter silent.


local Export = Export or {}

local Logger = require("BSE-DCS-Export/logger")
local DCS = DCS or {}

-- Main entry point:

local BSE = {
    frameCounter = 0,
    relay_address = {
        ip = "127.0.0.1",
        port = 6666
    },

    last_updated = {
        playerId = 0,
        position = 0,
        worldObjects = {},
        mission = 0
    },

    max_send_ops = 256,

    sender = require("BSE-DCS-Export/udp")
    --sender = require("BSE-DCS-Export/tcp"),
}

function BSE:Start()
    Logger:info("Starting...")
    self.sender:init(self.relay_address.ip, self.relay_address.port)
    Logger:info("Started")
end


function BSE:Stop()
    if self.sender ~= nil then
        self.sender:close()
    end

    Logger:info("closed.")
end

function BSE:shouldUpdate(last_updated, threshold)
    return self.frameCounter - last_updated > threshold
end

function BSE:UpdatePlayerUnit(threshold)
    if not self:shouldUpdate(self.last_updated.playerId, threshold) then return end
    
    self.sender:send({
        playerId = DCS.getPlayerUnit()                
    })   
    self.last_updated.playerId = self.frameCounter
end

function BSE:UpdatePlayerPosition(threshold)
    if not self:shouldUpdate(self.last_updated.position, threshold) then return end
    self.sender:send({
        position = Export.LoGetSelfData()                
    })
    self.last_updated.position = self.frameCounter
end

function BSE:UpdateWorldObjects(threshold)    
    local worldObjects = Export.LoGetWorldObjects()
    
    for id, unit in pairs(worldObjects) do
        local last_updated = self.last_updated.worldObjects[id] or 0
        if self:shouldUpdate(last_updated, threshold) then
            self.sender:send({
                worldObjects = {
                    [id] = unit
                }
            })
            self.last_updated.worldObjects[id] = self.frameCounter
        end

        if self.sender.sent_objects > self.max_send_ops then 
            break
        end    
    end
end

function BSE:UpdateMissionData(threshold)
    if not self:shouldUpdate(self.last_updated.mission, threshold) then return end    
    local mission = DCS.getCurrentMission()
    self.sender:send(DCS.getCurrentMission())
    self.last_updated.mission = self.frameCounter
end

function BSE:Update()
    self.frameCounter = self.frameCounter + 1

    Logger:debug("Updating..")
  
    if self.frameCounter == 1 then
        self:UpdatePlayerUnit(0)
        self:UpdatePlayerPosition(0)
        self:UpdateWorldObjects(0)
        self:UpdateMissionData(0)    
        Logger:debug("Early update, at first...")
        return
    end

    self:UpdatePlayerUnit(200)
    self:UpdatePlayerPosition(10)
    self:UpdateWorldObjects(30)
    self:UpdateMissionData(1024)

    Logger:debug("Updated.")
end

DCS.setUserCallbacks({
    onSimulationStart = function() BSE:Start() end,
    onSimulationStop = function() BSE:Stop() end,
    onSimulationFrame = function() BSE:Update() end
})
