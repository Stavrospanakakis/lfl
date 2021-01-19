export default function Section(props) {
    return (
        <div className="mt-28" id={props.id}>
            <div className="font-bold text-2xl lg:text-4xl">{props.title}</div>
            <div className="text-lg lg:text-xl italic text-gray-600">{props.subtitle}</div>
            <div className="">{props.description}</div>
            {props.children}
        </div>
    )
}