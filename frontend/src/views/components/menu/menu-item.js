import { Div, MaterialIcon, Text } from "../../elements/index.js";

export class MenuItem {
    constructor({ title, iconName }) {
        return new Div({
            className: 'menu-item',
            style: {
                display: 'flex',
                flexDirection: 'row',
                width: '100%',
                justifyContent: 'baseline',
                margin:'auto',
                textAlign: 'center'
            },
            children: [
                new MaterialIcon({
                    iconName: iconName,
                    style: {
                        color: 'var(--bs-blue)',
                        marginRight: '10px',
                        verticalAlign:'baseline',
                    }
                }),
                new Div({
                    className: 'menu-item-title',
                    style: {
                        color: 'var(--bs-blue)',
                        fontWeight: 'bold',
                    },
                    children: [new Text({ text: title })]
                })
            ]
        })
    }
}