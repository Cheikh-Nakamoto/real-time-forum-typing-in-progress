import { post } from "../../api/api.js";
import { goTo } from "../../lib/pager/navigation.js";
import { getView } from "../../lib/lib.js";
import { Div,TextField,Button,Text } from "../elements/index.js"

export class LoginPanel {
    constructor() {
        return new Div({
            id: "loginPanel",
            className: "panel-transition",
            style: {
                display: "none",
                flexDirection: "column",
                justifyContent: "center",   
            },
            children: [
                new Div({
                    style: {
                        display: "flex",
                        flexDirection: "column",
                        justifyContent: "center",
                        alignItems: "center",
                        gap: "2rem",
                        padding: "2rem",
                    },
                    children: [
                        new TextField({
                            id: "login",
                            className: "auth-field",
                            placeholder: "username/email",
                            type: "text",
                        }),
                        new TextField({
                            id: "loginPassword",
                            className: "auth-field",
                            placeholder: "password",
                            type: "password",
                        })
                    ]
                }),
                new Button({
                    id: "loginButton",
                    className: "auth-button",
                    children: [
                        new Text({ text: "Login" })
                    ],
                    listeners:{
                        onclick:()=>{
                            post("/auth/signin",this.formData).then((response)=>{
                                if (response) {
                                    localStorage.setItem("auth-token",response.authToken)
                                    const event = new CustomEvent("logged",{detail:{user:response.user}})
                                    dispatchEvent(event)
                                    goTo("contentPage")
                                }
                            }).catch((error)=> console.log(error))
                        }
                    }
                })
            ]
        })
    }

    get formData(){
        return {
            identifier: getView('login').element.value,
            password: getView('loginPassword').element.value,
        };
    }
}
