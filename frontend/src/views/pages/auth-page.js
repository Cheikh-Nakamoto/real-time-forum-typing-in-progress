import { getView, setPage } from "../../lib/lib.js";
import { Divider } from "../components/divider.js";
import { LoginPanel, SignupPanel, Snackbar } from "../components/index.js";
import { Div, Text } from "../elements/index.js";
import { AuthLayout } from "../layout/auth-layout.js";


export class AuthPage {
    constructor(options) {
        this.id = "auth"
        this.pagerId = options.pagerId
        this.title = "Authentication"
        this.loginPanel = getView("loginPanel") || new LoginPanel()
        this.signupPanel = getView("signupPanel") || new SignupPanel()
        this.currentPanel = location.pathname === "/auth/signin" ? this.loginPanel : this.signupPanel
        if (/\/auth\/\w+/.test(location.pathname)) {
            this.currentPanel.element.style.display = "flex"
        }

        setPage(this)
    }
    get element() {
        return new AuthLayout({
            style: {
                width: '100M',
                height: '100vh',
            },
            children: [
                new Div({
                    className: "auth-page",
                    style: {
                        display: "flex",
                        flexDirection: "column",
                        justifyContent: "center",
                        alignItems: "center",
                        paddingTop: "1rem",
                        height: "100%",
                        gap: "1rem",
                    },
                    children: [
                        new Div({
                            style: {
                                width: "45%",
                                backgroundColor: "white",
                                borderRadius: "5px",
                                paddingBottom: "2rem",
                                display: "flex",
                                flexDirection: "column",
                                justifyContent: "center",
                                boxShadow: "0px 10px 15px -3px rgba(0,0,0,0.1)",
                                maxWidth: "520px",
                            },
                            children: [
                                new Div({
                                    id: "form-wrapper",
                                    style: {
                                        display: "flex",
                                        flexDirection: "row",
                                        // justifyContent: "space-between",
                                        alignItems: "center",
                                        // gap: "1rem"
                                    },
                                    children: [
                                        new Div({
                                            id: "login-label",
                                            className: `auth-label ${this.currentPanel.id === "loginPanel" && "active"}`,
                                            style: {
                                                backgroundColor: this.currentPanel.id === "loginPanel" && "var(--bs-blue)",
                                                flex: 2,
                                                textAlign: "center",
                                                padding: "1rem"
                                            }, children: [new Text({ text: "Login" })],
                                            listeners: {
                                                onclick: () => {
                                                    getView("register-label").element.classList.remove("active")
                                                    getView("login-label").element.classList.add("active")
                                                    this.switchPanel("loginPanel")
                                                }
                                            }
                                        }),
                                        new Divider({
                                            width: "1px",
                                            height: "100%",
                                        }),
                                        new Div({
                                            id: "register-label",
                                            className: `auth-label ${this.currentPanel.id === "signupPanel" && "active"}`,
                                            style: {
                                                flex: 2,
                                                textAlign: "center",
                                                padding: "1rem",
                                                backgroundColor: this.currentPanel.id === "signupPanel" && "var(--bs-blue)"
                                            }, children: [new Text({ text: "Register" })],
                                            listeners: {
                                                onclick: () => {
                                                    getView("register-label").element.classList.add("active")
                                                    getView("login-label").element.classList.remove("active")
                                                    this.switchPanel("signupPanel")
                                                }
                                            }
                                        })
                                    ]
                                }),
                                new Divider({
                                    width: "100%",
                                }),
                                this.loginPanel,
                                this.signupPanel
                            ]
                        })
                    ]
                }),
                new Snackbar()
            ]
        }).element

    }

    get path() {
        return this.currentPanel === this.loginPanel ? "/auth/signin" : "/auth/signup"

    }

    switchPanel(panelId) {
        if (panelId !== this.currentPanel.id) {
            const view = getView(panelId)
            this.currentPanel.element.style.display = "none";
            view.element.style.display = "flex";
            this.currentPanel = view
            window.history.replaceState("", this.title, this.path);
        }
    }

    setDefaultPanel() {
        if (this.currentPanel === this.loginPanel) {
            const label = getView("login-label").element
            label.classList.add("active")

        } else {
            const label = getView("register-label").element
            label.classList.add("active")
        }
    }
}