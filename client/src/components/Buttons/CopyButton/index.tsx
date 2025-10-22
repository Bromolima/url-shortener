import { Copy } from 'lucide-react'
import styles from './styles.module.css'

type CopyButtonProps = { }& React.ComponentProps<'button'>

export function CopyButton({...props}:CopyButtonProps) {
    return (
        <button className={styles.button} {...props}>
            <Copy/>
        </button>
    )
}