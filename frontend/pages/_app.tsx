import '../styles/globals.css'
import type { AppProps } from 'next/app'
import {Col, Container, Navbar, Row} from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';

function SkyNet({ Component, pageProps }: AppProps) {
  return <>
    <Navbar bg="danger" sticky="top">
      <Container>
        <Navbar.Brand>SkyNET</Navbar.Brand>
      </Container>
    </Navbar>
    <Component {...pageProps} />
  </>
}
export default SkyNet
