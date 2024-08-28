import { post, setWSConnection } from "../../api/api.js";
import { getView } from "../../lib/lib.js";
import { goTo } from "../../lib/pager/navigation.js";
import { Div, TextField, Button, Text, Form, RadioButton } from "../elements/index.js"

export class SignupPanel {
    constructor() {
        return new Form({
            id: "signupPanel",
            className: "panel-transition",
            style: {
                display: "none",
                flexDirection: "column",
                justifyContent: "center",
                // transition: "visibility 0.7s, left 0.7s",
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
                        new Div({
                            style: {
                                display: "flex",
                                width: "100%",
                                flexDirection: "row",
                                justifyContent: "space-between",
                                alignItems: "center",
                                gap: "1rem"
                            },
                            children: [
                                new TextField({
                                    id: "fname",
                                    className: "auth-field",
                                    placeholder: "firstname",
                                    type: "text",
                                }),
                                new TextField({
                                    id: "lname",
                                    className: "auth-field",
                                    placeholder: "lastname",
                                    type: "text",
                                }),
                            ]
                        }),
                        new Div({
                            style: {
                                display: "flex",
                                width: "100%",
                                flexDirection: "row",
                                justifyContent: "space-between",
                                gap: "1rem"
                            },
                            children: [
                                new Div({
                                    style: {
                                        display:"flex",
                                        flexDirection: "row",
                                        justifyContent: "space-between",
                                        alignItems: "center",
                                        gap: "1rem"
                                    },
                                    children: [
                                        new Text({
                                            text: "Age:",
                                        }),
                                        new TextField({
                                            type: "number",
                                            name: "age",
                                            id: "age",
                                            className: "auth-field",
                                            placeholder: "age",
                                            style:{
                                                width: "64px"
                                            },
                                            attr: {
                                                min: "13",
                                                max: "99",
                                            }

                                        }),
                                    ]
                                }),
                                new Div({
                                    style: {
                                        display:"flex",
                                        flexDirection: "row",
                                        justifyContent: "space-between",
                                        alignItems: "center",
                                        gap: "1rem"
                                    },
                                    children: [
                                        new Text({
                                            text: "Gender:",
                                        }),
                                        new Div({
                                            children:[
                                                new Text({
                                                    text: "male",
                                                }),
                                                new RadioButton({
                                                    name: "gender",
                                                    id: "male",
                                                    attr:{
                                                        value:"male"
                                                    }
                                                    // className: "auth-field",
                                                }),
                                            ]
                                        }),
                                        new Div({
                                            children:[
                                                new Text({
                                                    text: "female",
                                                }),
                                                new RadioButton({
                                                    name: "gender",
                                                    id: "female",
                                                    attr:{
                                                        value:"female"
                                                    }
                                                }),
                                            ]
                                        }),
                                    ]
                                })
                            ]
                        })
                        ,
                        new TextField({
                            id: "uname",
                            className: "auth-field",
                            placeholder: "username",
                            type: "text",
                        }),

                        new TextField({
                            id: "email",
                            className: "auth-field",
                            placeholder: "email",
                            type: "email",
                        }),
                        new TextField({
                            id: "password",
                            className: "auth-field",
                            placeholder: "password",
                            type: "password",
                        }),
                        new TextField({
                            id: "rpass",
                            className: "auth-field",
                            placeholder: "repeat password",
                            type: "password",
                        }),
                    ]
                }),
                new Button({
                    id: "submitButton",
                    className: "auth-button",
                    children: [
                        new Text({ text: "Register" })
                    ]
                })
            ],
            listeners: {
                onsubmit: (event) => {
                    event.preventDefault()
                    post("/auth/signup", this.formData).then((response) => {
                        if (response) {
                            localStorage.setItem("auth-token", response.authToken)

                            const event = new CustomEvent("logged", { detail: { user: response.user } })
                            dispatchEvent(event)
                            goTo("contentPage")
                        }
                    }).catch((error) => console.log(error))

                }
            }
        })
    }

    get formData() {
        
        const checked = getView('signupPanel').element.querySelector('input[name=gender]:checked');
        return {
            firstname: getView('fname').element.value,
            lastname: getView('lname').element.value,
            username: getView('uname').element.value,
            email: getView('email').element.value,
            password: getView('password').element.value,
            age: getView('age').element.value,
            gender: checked && checked.value
        };

    }
    
}

