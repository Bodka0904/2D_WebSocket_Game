
class Setup {
    constructor() {
        this.gl
        this.canvas
        this.Sprite = {}
        this.Skin = {}
        this.Map = {}
        this.Item = {}
        this.Resource = {}
        this.Img = {}
        this.Loaded = false
        this.world = new World(this.Sprite.player)
      
    }
    InitGL() {
        this.canvas = document.querySelector("#glCanvas");
        this.gl = this.canvas.getContext("2d");
        this.gl.font = '10px Arial';
    }
    Load() {

        //Img
        this.Img.BuildMode = new Image()
        this.Img.BuildMode.onload = OnloadCallback
        this.Img.BuildMode.src = "/client/images/items/build_mode.png"

        //Sprites
        this.Sprite.player = new Image()
        this.Sprite.player.onload = OnloadCallback
        this.Sprite.player.src = "/client/images/sprites/player_sprite.png"
        /////////////////////////////////////
        this.world = new World(this.Sprite.player) 

        // Maps
        this.Map.field = new Image()
        this.Map.field.onload = OnloadCallback
        this.Map.field.src = "/client/images/maps/map.png"

        this.Map.forest = new Image()
        this.Map.forest.onload = OnloadCallback
        this.Map.forest.src = "/client/images/maps/map2.png"

        // Items
        this.Item.shield = new Image()
        this.Item.shield.onload = OnloadCallback
        this.Item.shield.src = "/client/images/items/shield.png"

        this.Item.staff = new Image()
        this.Item.staff.onload = OnloadCallback
        this.Item.staff.src = "/client/images/items/staff.png"

        this.Item.sword = new Image()
        this.Item.sword.onload = OnloadCallback
        this.Item.sword.src = "/client/images/items/sword.png"
        
        this.Item.wood = new Image()
        this.Item.wood.onload = OnloadCallback
        this.Item.wood.src = "/client/images/material/wood.png"

        this.Item.gravel = new Image()
        this.Item.gravel.onload = OnloadCallback
        this.Item.gravel.src = "/client/images/material/gravel.png"

        this.Item.stone = new Image()
        this.Item.stone.onload = OnloadCallback
        this.Item.stone.src = "/client/images/material/stone.png"

        //Resources
        this.Resource.tree = new Image()
        this.Resource.tree.onload = OnloadCallback
        this.Resource.tree.src = "/client/images/sources/tree.png"

        this.Resource.stone = new Image()
        this.Resource.stone.onload = OnloadCallback
        this.Resource.stone.src = "/client/images/sources/stone.png"

       
    }
    


}







