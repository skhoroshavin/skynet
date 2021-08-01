import '../styles/globals.css'
import type { AppProps } from 'next/app'
import Link from "next/link";
import useSWR, {mutate} from "swr";
import auth from "../services/auth";
import React from "react";

const UserSignedOut = () => {
  return <>
    <Link href="/signin" passHref>
      <Nav.Link>Sign in</Nav.Link>
    </Link>
    <Nav.Link disabled>|</Nav.Link>
    <Link href="/signup" passHref>
      <Nav.Link>Sign up</Nav.Link>
    </Link>
  </>
}

const UserSignedIn = (props: {userId: string|null|undefined}) => {
  const signOut = async (e: React.MouseEvent<HTMLElement>) => {
    e.preventDefault()
    await auth.signOut()
    await mutate("me")
  }

  return <NavDropdown id="user-id" title={props.userId}>
    <NavDropdown.Item href="#/signout" onClick={signOut}>
      Sign out
    </NavDropdown.Item>
  </NavDropdown>
}

const UserInfo = (props: {userId: string|null|undefined}) => {
  if (!props.userId) {
    return <UserSignedOut/>
  }
  return <UserSignedIn {...props}/>
}

function SkyNet({ Component, pageProps }: AppProps) {
  const userId = useSWR("me", auth.me).data

  return <>
    <Navbar bg="primary" fixed="top">
      <Container>
        <Navbar.Brand>SkyNET</Navbar.Brand>
        <Form className="d-flex me-auto">
          <Form.Control placeholder="Search"/>
          <Button variant="outline-success">Search</Button>
        </Form>
        <Nav>
          <UserInfo userId={userId}/>
        </Nav>
      </Container>
    </Navbar>
    <div style={{paddingTop: "60px"}}>
      <Component {...pageProps}/>
    </div>
  </>
}
export default SkyNet
