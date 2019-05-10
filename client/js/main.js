function Init(){

    setup.InitGL()
    setup.Load()
}

function GameLoop(){
   
    AddItemToInventory(setup.Item.staff,3)
  
    setInterval(function () {
        if (setup.Loaded){
    
        setup.world.DrawWorld(setup.gl,setup.Map.field,setup.Item,setup.Resource,setup.Img)
        
    }
    }, 15);
    
}

