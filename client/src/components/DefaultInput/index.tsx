import styles from  './styles.module.css'

type DefaultInputProps = { 
    id: string
} & React.ComponentProps<'input'>

export function DefaultInput({id, type, ...props}:DefaultInputProps){
    return (
        <div className={styles.container}>
            <label htmlFor={id}></label>
            <input className={styles.input} id={id} type={type} {...props}/>
        </div>
    )
}
