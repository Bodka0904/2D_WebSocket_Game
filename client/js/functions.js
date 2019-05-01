
const role = {
    man:{attack:10,intellect:8,defense:12},
    goblin:{attack:14,intellect:5,defense:10},
    warrior:{attack:12,intellect:8,defense:10},
    wizard:{attack:6,intellect:15,defense:7}
}


// Function that allows check only one checkbox with name check and shows attributes of checked player
function checkOnlyOne(checkbox) {
    var checkboxes = document.getElementsByName('check')
    checkboxes.forEach((item) => {
        if (item !== checkbox) item.checked = false

        if (item.checked == true) {
            if(item.value == "man"){
                document.getElementById("attack").innerText = role.man.attack
                document.getElementById("intellect").innerText = role.man.intellect
                document.getElementById("defense").innerText = role.man.defense
            }
            if(item.value == "goblin"){
                document.getElementById("attack").innerText = role.goblin.attack
                document.getElementById("intellect").innerText = role.goblin.intellect
                document.getElementById("defense").innerText = role.goblin.defense
            }
            if(item.value == "warrior"){
                document.getElementById("attack").innerText = role.warrior.attack
                document.getElementById("intellect").innerText = role.warrior.intellect
                document.getElementById("defense").innerText = role.warrior.defense
            }
            if(item.value == "wizard"){
                document.getElementById("attack").innerText = role.wizard.attack
                document.getElementById("intellect").innerText = role.wizard.intellect
                document.getElementById("defense").innerText = role.wizard.defense
            }
        }
    })

}


function AddItemToInventory(skin,slotNum){
    document.getElementById("slot"+slotNum).src = skin.src
}

function DeleteItemFromInventory(slotNum){
    document.getElementById("slot"+slotNum).src = "/client/images/items/emptySlot.png"
}

