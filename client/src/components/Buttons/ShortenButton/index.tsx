import styles from './styles.module.css'

type ShortenButtonProps = { }& React.ComponentProps<'button'>

export function ShortenButton({...props}:ShortenButtonProps) {
    return (
        <button className={styles.button} {...props}>
            Encurtar URL
        </button>
    )
}