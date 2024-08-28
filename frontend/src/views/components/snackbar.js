import { Div } from "../elements/index.js"
import { getView } from "../../lib/lib.js"
export class Snackbar {
    constructor(){
        addEventListener("errorOccur",(event)=>{
            const error = event.detail
            const msgDiv = getView("errorMsg").element

            msgDiv.innerHTML = error.msg

            this.show()
        })
    }

    get element(){
        return new Div({
            id:"snackbar",
            className:"snackbar",
            
            children:[  
                new Div({
                    className: "snackbar-message",
                    children:[
                        new Div({
                            id:"errorMsg",
                        })
                    ]
                })
            ]
        }).element
    }

    show() {
        document.getElementById("snackbar").classList.add('show')

        setTimeout(() => {
            document.getElementById("snackbar").classList.remove('show')
        }, 3000);
    }
}