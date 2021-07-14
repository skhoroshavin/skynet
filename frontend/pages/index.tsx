import Link from "next/link";

export default function Home() {
  return (
      <>
          <div>Welcome to SkyNET, social network for cyborgs and all that stuff!</div>
          <Link href="/signup">Signup</Link>
      </>
  )
}
