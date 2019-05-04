
class Player {
    constructor(name, PosX, PosY, Sprite_sheet, ID, HP, Inventory) {
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
            Down:false
        }
        this.Sprite = new Sprite(Sprite_sheet,175,28,35)

    }

    UpdateData(NewX, NewY,Control) {
        this.PosX = NewX
        this.PosY = NewY

        this.Control.Up = Control.Up
        this.Control.Down = Control.Down

        if (Control.Right == true)
        {
            this.Control.Right = true
            this.Control.Left = false
        }
        if (Control.Left == true){

            this.Control.Left = true
            this.Control.Right = false
        }

    }
   
    Draw(gl) {
        gl.fillText(this.ID, this.PosX, this.PosY - 1)
        //gl.drawImage(this.Skin, this.PosX, this.PosY)
        //gl.drawImage(this.Skin,0,0,20,10,this.PosX,this.PosY,35,28)

        
        if (this.Control.Right){
            this.Sprite.Animate(gl,this.PosX,this.PosY,0,0,this.Control.Right)
        }
        if (this.Control.Left){
            this.Sprite.Animate(gl,this.PosX,this.PosY,1,4,this.Control.Left)
        }else {
            this.Sprite.Animate(gl,this.PosX,this.PosY,0,0,this.Control.Right)
        }
        

    }

}


