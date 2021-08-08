import {user, UserData} from "../../api/user";
import {UserCard} from "../../components/user-card";
import {Card} from "../../components/common/card";
import {ChipIcon} from "@heroicons/react/solid";

export const getServerSideProps = async (context: any) => {
    let res = await user.data(context.params.userId)
    if (res instanceof Error) {
        return {notFound: true}
    }

    return {props: res}
}

export default function UserPage(props: UserData) {
    return <>
        <Card className="w-1/3 text-grey-200">
            <ChipIcon/>
        </Card>
        <UserCard className="w-2/3 ml-4" user={props}/>
    </>
}
