import {UserData} from "../api/user";
import {Card, CardDivider, CardTitle} from "./common/card";

type InfoRowProps = {
    label: string,
    value: any
}

const UserDataRow = (props: InfoRowProps) => {
    if (!props.value)
        return null

    return <div className="flex mt-2">
        <div className="w-2/5 text-grey-500">{props.label}</div>
        <div className="flex-1">{props.value}</div>
    </div>
}

type UserCardProps = {
    user: UserData,
    className?: string
}

export const UserCard = ({ user, className }: UserCardProps) => {
    const birthday = user.birthday && new Date(user.birthday).toDateString()

    return <Card className={className}>
        <CardTitle>{user.firstName} {user.lastName}</CardTitle>
        <CardDivider className="mt-4"/>
        <div className="mt-4">
            <UserDataRow label="Birthday:" value={birthday}/>
            <UserDataRow label="Homecity:" value={user.city}/>
            <UserDataRow label="Interests:" value={user.interests}/>
        </div>
    </Card>
}
