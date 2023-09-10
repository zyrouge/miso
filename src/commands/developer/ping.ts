import { TofuCommand } from "@/core/command";
import { emojis } from "@/utils/emojis";
import { ErisUtils } from "@/utils/eris";

export const pingCommand: TofuCommand = {
    config: {
        name: "ping",
        description: "Ping pong!",
    },
    invoke: async (_, __) => {
        return {
            message: ErisUtils.prettyMessage(emojis.pingPong, "Pong!"),
        };
    },
};
