class Setup {
    constructor() {
        this.gl
        this.canvas
        this.Players
        this.Skin = {}
        this.Map = {}
        this.player_list = []
    }
    InitGL() {
        this.canvas = document.querySelector("#glCanvas");
        this.gl = this.canvas.getContext("2d");

        this.gl.font = '10px Arial';
    }
    Load() {

        // Skins
        this.Skin.player = new Image()
        this.Skin.player.src = "/client/images/player.png"

        this.Skin.bat = new Image()
        this.Skin.bat.src = "/client/images/bat.png"

        this.Skin.warrior = new Image()
        this.Skin.warrior.src = "/client/images/warrior.png"


        // Maps
        this.Map.field = new Image()
        this.Map.field.src = "/client/images/map.png"

        this.Map.forest = new Image()
        this.Map.forest.src = "/client/images/map2.png"

    }

    DrawMap() {
        this.gl.drawImage(this.Map.forest, 0, 0)
    }

    AddPlayer(serverData) {
        
        if (this.player_list.length == 0) {
            //Add first player - client
            this.player_list.push(new Player("", 250, 250, "", serverData[serverData.length - 1].ID))
            console.log("New player added")

            //Add all players that were connected earlier
            for (var i = serverData.length - 2; i >= 0; i--) {
                this.player_list.push(new Player("", 250, 250, "", serverData[i].ID))
                console.log("New player added")
            }

        } else if (serverData.length > this.player_list.length) {
            // Wait some time 
            setTimeout(function () {
            }, 50)

            //Add new connected players
            this.player_list.push(new Player("", 250, 250, "", serverData[serverData.length - 1].ID))
            console.log("New player added")

        }

    }


    DeletePlayer(serverData) {

        if (serverData.length < this.player_list.length) {
            for (var j = 0; j < this.player_list.length; j++) {
                for (var i = 0; i < serverData.length; i++) {
                    if (this.player_list[j].ID != serverData[i].ID) {

                        this.player_list.splice(j, 1)
                        console.log("Player deleted")
                    }

                }
            }
        }

    }


}







