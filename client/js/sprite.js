class Sprite {
    constructor(sprite,width,height,columns,rows,delay){
        
        this.sprite = sprite
        this.width = width
        this.height = height
        this.frame_width = width / columns
        this.frame_height = height / rows
        this.rows = rows
        this.columns = columns
        
        this.frames = this.Init()
        this.startTime = 0
        this.delay = delay
        
    }
    Init(){
        var tmpPairs = []
        for (var j = 0; j < this.rows ; j++){
            for (var i = 0 ; i < this.columns ; i++){
                tmpPairs.push([i * this.frame_width,j * this.frame_height])
            }
        }
      
        return tmpPairs

    
    }


    
    Animate(gl,PosX,PosY,default_frame,next_frame,dt) {
            
            this.startTime += dt
            
            if (this.startTime <= this.delay/2){
                gl.drawImage(this.sprite,default_frame[0],default_frame[1],this.frame_width,this.frame_height,PosX,PosY,this.frame_width,this.frame_height)
            } else {
                gl.drawImage(this.sprite,next_frame[0],next_frame[1],this.frame_width,this.frame_height,PosX,PosY,this.frame_width,this.frame_height)
            }
            if (this.startTime == this.delay){
                this.startTime = 0
            }

    }



    
}