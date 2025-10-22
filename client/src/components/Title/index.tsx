import { Link2 } from 'lucide-react'
import styles from './styles.module.css'

export function Title() {
    return (
        <div className={styles.container}>
            <Link2 size={64}/>

            <h1>Encurte seus links</h1>

            <h6>Transforme URLs longas em links limpos, curtos e compartilháveis  em segundos</h6>
        </div>
    )
}