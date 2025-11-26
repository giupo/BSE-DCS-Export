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
    every_framecount = 30,

    sender = require("BSE-DCS-Export/udp")
    -- sender = require("BSE-DCS-Export/tcp")
}

local DcsWorldState = {
    playerId = 0,
    player_position = {},
    world_objects = {},
    mission_data = {},
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

function BSE:shouldUpdate(timeElapsed)
    return (self.frameCounter % self.every_framecount) == 0
end


function BSE:Update()
    self.frameCounter = self.frameCounter + 1
    Logger:debug("Updating..")
    if not self:shouldUpdate() then return end

    local new_state = {
        playerId = DCS.getPlayerUnit(),
        player_position = Export.LoGetSelfData(),
        world_objects = Export.LoGetWorldObjects(),
        mission_data = DCS.getCurrentMission()
    }

    -- local to_be_sent = tableDiff(DcsWorldState, new_state)    
    self.sender:send(new_state)

    DcsWorldState = new_state

    Logger:debug("Updated.")
end

DCS.setUserCallbacks({
    onSimulationStart = function() BSE:Start() end,
    onSimulationStop = function() BSE:Stop() end,
    onSimulationFrame = function() BSE:Update() end
})
