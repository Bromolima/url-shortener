import { CopyButton } from '../Buttons/CopyButton';
import styles from './styles.module.css';

interface ShortenedUrlCardProps {
    url: string;
}

export function ShortenedUrlCard({ url }: ShortenedUrlCardProps) {
    const handleCopy = () => {
        navigator.clipboard.writeText(url);
    };

    return (
        <div className={styles.card}>
            <p className={styles.url}>{url}</p>
            <CopyButton 
                onClick={handleCopy}
            />
        </div>
    );
}
