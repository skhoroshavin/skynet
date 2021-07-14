import '../styles/globals.css'
import type { AppProps } from 'next/app'

function SkyNet({ Component, pageProps }: AppProps) {
  return <Component {...pageProps} />
}
export default SkyNet
