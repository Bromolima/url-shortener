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
            <a className={styles.link}>{shortenedUrl}</a>
            <CopyButton 
                onClick={handleCopy}
            />
        </div>
    )
}