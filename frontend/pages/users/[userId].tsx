import {Card, Col, Container, Row} from "react-bootstrap";
import user, {UserData} from "../../services/user";
import {UserDataCard} from "../../components/user-data-card";

export const getServerSideProps = async (context: any) => {
    let res = await user.data(context.params.userId)
    if (res instanceof Error) {
        return {notFound: true}
    }

    return {props: res}
}

export default function UserPage(props: UserData) {
    return (
        <Container>
            <Row>
                <Col xs="12" md="4">
                    <Card style={{marginTop: 25}}>
                        <Card.Img src="/nuclear_blast.png"/>
                    </Card>
                </Col>
                <Col xs="12" md="8">
                    <UserDataCard {...props}/>
                </Col>
            </Row>
        </Container>
    )
}
