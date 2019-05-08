
class Player {
    constructor(name, PosX, PosY, Sprite_sheet, ID, HP,Energy,Inventory) {
        this.ID = ID
        this.HP = HP
        this.name = name
        this.PosX = PosX
        this.PosY = PosY
        this.Inventory = Inventory
        this.Control = {
            Right:false,
            Left:false,
            Up:false,
            Down:false,

            Action:{
                Attack:false,
                Mine:false,
                Build:false,
            }
        }
        this.Rotate = {
            Right:true,
            Left:false
        }
        this.Energy = Energy
        this.Sprite = new Sprite(Sprite_sheet,240,90,8,3,100)


    }

    UpdateData(serverData) {
        this.PosX = serverData.Position.X
        this.PosY = serverData.Position.Y

        this.Control.Up = serverData.Control.Up
        this.Control.Down = serverData.Control.Down
        this.Control.Right = serverData.Control.Right
        this.Control.Left = serverData.Control.Left

        this.Control.Action.Attack = serverData.Control.Action.Attack
        this.Control.Action.Mine = serverData.Control.Action.Mine
        this.Control.Action.Build = serverData.Control.Action.Build

        this.Inventory = serverData.Inventory
        

        if (serverData.Control.Right){
            this.Rotate.Right = true
            this.Rotate.Left = false
        }
        if (serverData.Control.Left){
            this.Rotate.Left = true
            this.Rotate.Right = false
        }

        if (serverData.Control.Action.Attack){
            this.Control.Action.Mine = false
            this.Control.Action.Build = false
        }
        if (serverData.Control.Action.Mine){
            this.Control.Action.Attack = false
            this.Control.Action.Build = false 
        }
        if (serverData.Control.Action.Build){
            this.Control.Action.Attack = false
            this.Control.Action.Mine = false  
        }


    }
    
   
    Draw(gl) {
        gl.fillText(this.ID, this.PosX, this.PosY - 1)
        //gl.drawImage(this.Skin, this.PosX, this.PosY)
        //gl.drawImage(this.Skin,0,0,20,10,this.PosX,this.PosY,35,28)

        
        if (this.Control.Right){
            this.Sprite.Animate(gl,this.PosX,this.PosY,this.Sprite.frames[8],this.Sprite.frames[9],2)
        }
        else if (this.Control.Left){
            this.Sprite.Animate(gl,this.PosX,this.PosY,this.Sprite.frames[11],this.Sprite.frames[10],2)
        }
        else if (this.Control.Up){
            this.Sprite.Animate(gl,this.PosX,this.PosY,this.Sprite.frames[6],this.Sprite.frames[7],2)
        }
        else if (this.Control.Down){
            this.Sprite.Animate(gl,this.PosX,this.PosY,this.Sprite.frames[4],this.Sprite.frames[5],2)
        }
        else if (this.Control.Action.Mine) {
            if (this.Rotate.Left){
                this.Sprite.Animate(gl,this.PosX,this.PosY,this.Sprite.frames[19],this.Sprite.frames[18],2)
            } else if (this.Rotate.Right){
                this.Sprite.Animate(gl,this.PosX,this.PosY,this.Sprite.frames[16],this.Sprite.frames[17],2)
            }
        }
        else if (this.Control.Action.Attack){
            if (this.Rotate.Left){
                this.Sprite.Animate(gl,this.PosX,this.PosY,this.Sprite.frames[15],this.Sprite.frames[14],2)
            } else if (this.Rotate.Right){
                this.Sprite.Animate(gl,this.PosX,this.PosY,this.Sprite.frames[12],this.Sprite.frames[13],2)
            }
        }
        
        
        
        else {
            this.Sprite.Animate(gl,this.PosX,this.PosY,this.Sprite.frames[0],this.Sprite.frames[1],2)
        }

    }

}


