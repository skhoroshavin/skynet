import user, {UserData} from "../../services/user";
import {UserDataCard} from "../../components/user-data-card";
import Image from "next/image"

export const getServerSideProps = async (context: any) => {
    let res = await user.data(context.params.userId)
    if (res instanceof Error) {
        return {notFound: true}
    }

    return {props: res}
}

export default function UserPage(props: UserData) {
    return <div>
         <Image src="/nuclear_blast.png" layout="fill" alt="Avatar"/>
         <UserDataCard {...props}/>
    </div>
}
