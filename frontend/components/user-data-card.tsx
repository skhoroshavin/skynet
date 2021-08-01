import {UserData} from "../services/user";

type InfoRowProps = {
    label: string,
    value: any
}

const UserDataRow = (props: InfoRowProps) => {
    if (!props.value)
        return null

    return <div>
        <div>{props.label}</div>
        <div>{props.value}</div>
    </div>
}

export const UserDataCard = (props: UserData) => {
    const birthday = props.birthday && new Date(props.birthday).toDateString()

    return <div>
        <div>{props.firstName} {props.lastName}</div>
        <hr/>
        <div>
            <UserDataRow label="Birthday:" value={birthday}/>
            <UserDataRow label="Homecity:" value={props.city}/>
            <UserDataRow label="Interests:" value={props.interests}/>
        </div>
    </div>
}
