import '../styles/globals.css'
import type { AppProps } from 'next/app'
import React from "react";
import { AppLayout } from '../components/app-layout';

function App({ Component, pageProps }: AppProps) {
    return <AppLayout>
        <Component {...pageProps}/>
    </AppLayout>
}
export default App
