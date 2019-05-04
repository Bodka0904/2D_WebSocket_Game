function Init(){

    setup.InitGL()
    setup.Load()
}

function GameLoop(){
   
    AddItemToInventory(setup.Item.staff,3)
  
    setInterval(function () {
        if (setup.Loaded){
    
        // Clear Canvas
        setup.gl.clearRect(0, 0, 500, 500)
        setup.DrawMap()
        

        if (serverData != null && setup.player_list.length != 0) {

            for (var i = 0; i < setup.player_list.length; i++) {
              
                setup.player_list[i].Draw(setup.gl)
            }
        }
    }
    }, 15);
    
}

