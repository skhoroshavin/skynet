import {UserData} from "../services/user";

type InfoRowProps = {
    label: string,
    value: any
}

const UserDataRow = (props: InfoRowProps) => {
    if (!props.value)
        return null

    return <Row>
        <Col xs="4" style={{color: "gray"}}>{props.label}</Col>
        <Col>{props.value}</Col>
    </Row>
}

export const UserDataCard = (props: UserData) => {
    const birthday = props.birthday && new Date(props.birthday).toDateString()

    return <Card body style={{marginTop: 25}}>
        <Card.Title>{props.firstName} {props.lastName}</Card.Title>
        <hr/>
        <Card.Text>
            <UserDataRow label="Birthday:" value={birthday}/>
            <UserDataRow label="Homecity:" value={props.city}/>
            <UserDataRow label="Interests:" value={props.interests}/>
        </Card.Text>
    </Card>
}
