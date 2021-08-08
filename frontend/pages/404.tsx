import {NuclearBlast} from "../components/common/nuclear-blast";

export default function NotFound() {
    return <>
        <NuclearBlast className="absolute fullscreen inset-x-0 text-primary-200 w-full h-full mx-auto p-6"/>

        <div className="fullscreen relative flex flex-col items-center justify-center">
            <div className="text-4xl text-primary-900 text-center font-mono font-bold">404 Not Found</div>
            <div className="h-1/3"/>
        </div>
    </>
}
