
class Player {
    constructor(name, PosX, PosY, skin, ID) {
        this.ID = ID
        this.name = name
        this.Skin = skin
        this.PosX = PosX
        this.PosY = PosY



        this.width = skin.width
        this.height = skin.height
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


