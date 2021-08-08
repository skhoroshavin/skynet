import useSWR, {mutate} from "swr";
import auth from "../client/auth";

const _SWR_CURRENT_USER = "/auth/me"

export const useCurrentUser = () => {
    const { data } = useSWR(_SWR_CURRENT_USER, auth.me)

    return data;
}
