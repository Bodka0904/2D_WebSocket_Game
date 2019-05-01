

class Setup {
    constructor() {
        this.gl
        this.canvas
        this.Players
        this.Skin = {}
        this.Map = {}
        this.Item = {}
        this.player_list = []
    }
    InitGL() {
        this.canvas = document.querySelector("#glCanvas");
        this.gl = this.canvas.getContext("2d");

        this.gl.font = '10px Arial';
    }
    Load() {

        // Skins
        this.Skin.man = new Image()
        this.Skin.man.src = "/client/images/players/man.png"

        this.Skin.warrior = new Image()
        this.Skin.warrior.src = "/client/images/players/warrior.png"

        this.Skin.wizard = new Image()
        this.Skin.wizard.src = "/client/images/players/wizard.png"

        this.Skin.goblin = new Image()
        this.Skin.goblin.src = "/client/images/players/goblin.png"

        // Maps
        this.Map.field = new Image()
        this.Map.field.src = "/client/images/maps/map.png"

        this.Map.forest = new Image()
        this.Map.forest.src = "/client/images/maps/map2.png"

        // Items
        this.Item.shield = new Image()
        this.Item.shield.src = "/client/images/items/shield.png"

        this.Item.staff = new Image()
        this.Item.staff.src = "/client/images/items/staff.png"

        this.Item.sword = new Image()
        this.Item.sword.src = "/client/images/items/sword.png"
    }
    GetSkin(serverData){
        
        if (serverData == "goblin")
        {
          
            return this.Skin.goblin
            
        }
        if (serverData == "man")
        {
            
            return this.Skin.man
        }
        if (serverData == "wizard")
        {
         
            return this.Skin.wizard
        }
        if (serverData == "warrior")
        {
           
            return this.Skin.warrior
        }

        
    }


    DrawMap() {
        this.gl.drawImage(this.Map.field, 0, 0)
    }

    AddPlayer(serverData) {
     
        if (this.player_list.length == 0) {
            //Add first player - client
           
          
            this.player_list.push(new Player("", 250, 250, this.GetSkin(serverData[serverData.length - 1].Class), serverData[serverData.length - 1].ID,serverData[serverData.length - 1].Inventory))
            console.log("New player added")
            

            //Add all players that were connected earlier
            for (var i = serverData.length - 2; i >= 0; i--) {
                
                this.player_list.push(new Player("", 250, 250, this.GetSkin(serverData[i].Class), serverData[i].ID))
                console.log("New player added")
            }

        } else if (serverData.length > this.player_list.length) {
            // Wait some time 
            setTimeout(function () {
            }, 50)
            
            //Add new connected players
            this.player_list.push(new Player("", 250, 250, this.GetSkin(serverData[serverData.length - 1].Class), serverData[serverData.length - 1].ID))
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







