import { CopyButton } from '../Buttons/CopyButton';
import styles from './styles.module.css'

type UrlOutputProps = {
    shortenedUrl: string
}

export function UrlOutput({shortenedUrl} : UrlOutputProps) {
    const handleCopy = () => {
        navigator.clipboard.writeText(shortenedUrl);
    };

    return (
        <div className={styles.output}>
            <p className={styles.link}>{shortenedUrl}</p>
            <CopyButton 
                onClick={handleCopy}
            />
        </div>
    )
}