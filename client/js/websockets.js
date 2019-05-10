let setup = new Setup()


let serverData = null


//Function to get params from URL
function getUrlVars() {
    var vars = {};
    var parts = window.location.href.replace(/[?&]([^=#]+)=([^&#]*)/gi, function(m,key,value) {
        vars[key] = value;
    });
    return vars;
}     



ws.onopen = () => {

    console.log("Connected")
    var params = getUrlVars()
   
    //Send params from URL sent by server back to server to Init client player
    var X = parseFloat(params.PosX)
    var Y = parseFloat(params.PosY)
    
    ws.send(JSON.stringify({ID:params.ID,Position:{X:X,Y:Y}}))
   
    
}

ws.onmessage = (msg) => {

    if  (setup.Loaded){ 
        
        serverData = JSON.parse(msg.data)
   
        setup.world.UpdateData(serverData)

    
    }

}
//s

ws.onclose = () => {
    console.log("Disconnected")
}

// Control
document.onkeydown = function (event) {
    // Movement
    if (event.keyCode == 68) { //d
        ws.send(JSON.stringify({ Right: true }))
       
    }
    if (event.keyCode === 83){ //s
        ws.send(JSON.stringify({ Down: true }))
       
    }
    if (event.keyCode === 65){ //a
        ws.send(JSON.stringify({ Left: true }))
       
    }
    if (event.keyCode === 87) { //w
        ws.send(JSON.stringify({ Up: true }))
    
    }
    // Attack
    if(event.keyCode === 97)//1
        ws.send(JSON.stringify({Action:{Attack: true}}))
    if(event.keyCode === 98)//2
        ws.send(JSON.stringify({Action:{Mine: true}}))
    if(event.keyCode === 99)//3
        ws.send(JSON.stringify({Action:{Build: true}}))
    if (event.keyCode === 100) //4
        ws.send(JSON.stringify({Action:{Pick: true}}))

    
}
document.onkeyup = function (event) {
    // Movement
    if (event.keyCode == 68) { //d
        ws.send(JSON.stringify({ Right: false }))
    
    }
    if (event.keyCode === 83){ //s
        ws.send(JSON.stringify({ Down: false }))
        
    }
    if (event.keyCode === 65){ //a
        ws.send(JSON.stringify({ Left: false }))
      
    }
    if (event.keyCode === 87) { //w
        ws.send(JSON.stringify({ Up: false }))
      
    }
    // Attack
    if(event.keyCode === 97) //1
        ws.send(JSON.stringify({Action:{Attack: false}}))
    if(event.keyCode === 98)//2
        ws.send(JSON.stringify({Action:{Mine: false}}))
    if(event.keyCode === 99)//3
        ws.send(JSON.stringify({Action:{Build: false}}))
    if (event.keyCode === 100) //4
        ws.send(JSON.stringify({Action:{Pick: false}}))
}

window.ws = ws

