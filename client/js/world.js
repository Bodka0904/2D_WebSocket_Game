
class World {
    constructor(Sprite_sheet)
    {
        this.players = []
        this.MapSkin
        this.Sprite = new Sprite(Sprite_sheet,240,90,8,3,100)

    }

    UpdateData(serverData){
        if (serverData.length > this.players.length){
            this.players = serverData.slice(0)
        }
        for (var i = 0; i < serverData.length; i++){
            if (serverData[i].World.Creatures.length != this.players[i].World.Creatures.length){
                this.players = serverData.slice(0)
            } else if (serverData[i].World.Objects.length != this.players[i].World.Objects.length){
                this.players = serverData.slice(0)
            } else if (serverData[i].World.Items.length != this.players[i].World.Items.length){
                this.players = serverData.slice(0)
            } else {
                this.players[i].Position.X = serverData[i].Position.X
                this.players[i].Position.Y = serverData[i].Position.Y

                this.players[i].Face = serverData[i].Face
                this.players[i].BuildMode = serverData[i].BuildMode
                this.players[i].Control = serverData[i].Control

                for (var j = 0; j < this.players[i].World.Creatures[j]; j++){
                    this.players[i].World.Creatures.Position.X = serverData[i].World.Creatures[j].Position.X 
                    this.players[i].World.Creatures.Position.Y = serverData[i].World.Creatures[j].Position.Y 
                }
            }

            
        }
    }

    DrawWorld(gl,MapSkin,Items,Resources,Img){

        gl.clearRect(0, 0, 1000, 800)

        gl.drawImage(MapSkin,0,0)

        if (this.players.length != 0){
        for (var i = 0; i < this.players[0].World.Resources.length;i++){
            if (this.players[0].World.Resources[i].Name == "Tree") {
                gl.drawImage(Resources.tree,this.players[0].World.Resources[i].Position.X,this.players[0].World.Resources[i].Position.Y)
            }
            if (this.players[0].World.Resources[i].Name == "Stone") {
                gl.drawImage(Resources.stone,this.players[0].World.Resources[i].Position.X,this.players[0].World.Resources[i].Position.Y)
            }
        }

        
            for (var i = 0; i < this.players[0].World.Objects.length;i++){
                if (this.players[0].World.Objects[i].Name == "Wood") {
                    gl.drawImage(Items.wood,this.players[0].World.Objects[i].Position.X,this.players[0].World.Objects[i].Position.Y)
                }
                if (this.players[0].World.Objects[i].Name == "Gravel") {
                    gl.drawImage(Items.gravel,this.players[0].World.Objects[i].Position.X,this.players[0].World.Objects[i].Position.Y)
                }
                if (this.players[0].World.Objects[i].Name == "Stone") {
                    gl.drawImage(Items.stone,this.players[0].World.Objects[i].Position.X,this.players[0].World.Objects[i].Position.Y)
                }
            }
        }

        for (var i = 0; i < this.players.length; i++){
           
            gl.fillText(this.players[i].ID, this.players[i].Position.X, this.players[i].Position.Y - 1)
            

            if (this.players[i].Control.Right){
                this.Sprite.Animate(gl,this.players[i].Position.X,this.players[i].Position.Y,this.Sprite.frames[8],this.Sprite.frames[9],2)
            }
            else if (this.players[i].Control.Left){
                this.Sprite.Animate(gl,this.players[i].Position.X,this.players[i].Position.Y,this.Sprite.frames[11],this.Sprite.frames[10],2)
            }
            else if (this.players[i].Control.Up){
                this.Sprite.Animate(gl,this.players[i].Position.X,this.players[i].Position.Y,this.Sprite.frames[6],this.Sprite.frames[7],2)
            }
            else if (this.players[i].Control.Down){
                this.Sprite.Animate(gl,this.players[i].Position.X,this.players[i].Position.Y,this.Sprite.frames[4],this.Sprite.frames[5],2)
            }
            else if (this.players[i].Control.Action.Mine) {
                if (this.players[i].Face == "Left"){
                    this.Sprite.Animate(gl,this.players[i].Position.X,this.players[i].Position.Y,this.Sprite.frames[19],this.Sprite.frames[18],2)
                } else if (this.players[i].Face == "Right"){
                    this.Sprite.Animate(gl,this.players[i].Position.X,this.players[i].Position.Y,this.Sprite.frames[16],this.Sprite.frames[17],2)
                }
                else {
                    this.Sprite.Animate(gl,this.players[i].Position.X,this.players[i].Position.Y,this.Sprite.frames[16],this.Sprite.frames[17],2)
                }
            }
            else if (this.players[i].Control.Action.Attack){
                if (this.players[i].Face == "Left"){
                    this.Sprite.Animate(gl,this.players[i].Position.X,this.players[i].Position.Y,this.Sprite.frames[15],this.Sprite.frames[14],2)
                } else if (this.players[i].Face == "Right"){
                    this.Sprite.Animate(gl,this.players[i].Position.X,this.players[i].Position.Y,this.Sprite.frames[12],this.Sprite.frames[13],2)
                } else {
                    this.Sprite.Animate(gl,this.players[i].Position.X,this.players[i].Position.Y,this.Sprite.frames[12],this.Sprite.frames[13],2)
                }
            }
            
            
            
            else {
                this.Sprite.Animate(gl,this.players[i].Position.X,this.players[i].Position.Y,this.Sprite.frames[0],this.Sprite.frames[1],2)
            }
            
       
            if (this.players[i].BuildMode == true){
                
                if (this.players[i].Face == "Up"){
                    gl.drawImage(Img.BuildMode,this.players[i].Position.X + 5, this.players[i].Position.Y - 20)
                }
                if (this.players[i].Face == "Down"){
                    gl.drawImage(Img.BuildMode,this.players[i].Position.X + 5 , this.players[i].Position.Y + 35)
                }
                if (this.players[i].Face == "Right"){
                    gl.drawImage(Img.BuildMode,this.players[i].Position.X + 35 , this.players[i].Position.Y + 10)
                }
                if (this.players[i].Face == "Left"){
                    gl.drawImage(Img.BuildMode,this.players[i].Position.X - 20, this.players[i].Position.Y + 10)
                }
            }

        }
        
    }

    
}

