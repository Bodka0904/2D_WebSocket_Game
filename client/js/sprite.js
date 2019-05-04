class Sprite {
    constructor(sprite,width,height,frame_width){
        
        this.sprite = sprite
        this.width = width
        this.height = height
        this.frame_width = frame_width
        this.numFrames = width/frame_width
       
        
    }


    Animate(gl,PosX,PosY,default_frame,next_frame,bool) {

        if (!bool){
            gl.drawImage(this.sprite,default_frame * this.frame_width,0,this.frame_width,this.height,PosX,PosY,this.frame_width,this.height)
        } else {
            gl.drawImage(this.sprite,next_frame * this.frame_width,0,this.frame_width,this.height,PosX,PosY,this.frame_width,this.height)
        }

    }



    
}