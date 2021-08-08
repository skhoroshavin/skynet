import {PropsWithChildren} from "react";

type CardProps = {
    className?: string
}

export const Card = ({ className, children }: PropsWithChildren<CardProps>) => {
    return <div className={`p-4 bg-white rounded-lg shadow-md ${className}`}>
        {children}
    </div>
}

type CardTitleProps = {
    children?: React.ReactFragment,
    className?: string
}

export const CardTitle = ({ className, children }: CardTitleProps ) => {
    return <div className={`text-2xl ${className}`}>
        {children}
    </div>
}

type CardDividerProps = {
    className?: string
}

export const CardDivider = ({ className }: CardDividerProps ) => {
    return <hr className={`border-grey-300 ${className}`}/>
}
