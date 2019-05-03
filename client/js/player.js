
class Player {
    constructor(name, PosX, PosY, skin, ID, HP, Inventory) {
        this.ID = ID
        this.HP = HP
        this.name = name
        this.Skin = skin
        this.PosX = PosX
        this.PosY = PosY
        this.Inventory = Inventory

    }


    UpdateData(NewX, NewY) {
        this.PosX = NewX
        this.PosY = NewY

    }
   
    Draw(gl) {
        gl.fillText(this.ID, this.PosX, this.PosY - 1)
        gl.drawImage(this.Skin, this.PosX, this.PosY)
    }

}


