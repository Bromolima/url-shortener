import { DefaultInput } from '../DefaultInput'
import styles from './styles.module.css'
import { UrlOutput } from '../UrlOutput'
import { useLink } from '../../contexts/LinkContext/LinkContext'
import { ShortenButton } from '../Buttons/ShortenButton'

export function UrlShortenerForm() {
  const {
    originalUrl,
    setOriginalUrl,
    shortenedUrl,
    setShortenedUrl,
    loading,
    setLoading,
    error,
    setError,
  } = useLink()

  const handleShortenUrl = async () => {
    console.log('Valor de originalUrl ao clicar:', originalUrl)
    if (!originalUrl) return

    try {
      setLoading(true)
      setError('')

    const response = await fetch('http://localhost:8080/v1/shorten', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ url: originalUrl }),
    })

    if (!response.ok) {
      throw new Error('Falha ao encurtar a URL')
    }

    const data = await response.json()
    console.log('Resposta da API:', data)
    setShortenedUrl(data.short_code)
  } catch (err) {
    console.error('Erro na chamada da API:', err)
    setError('Erro ao encurtar URL')
  } finally {
    setLoading(false)
  }
}

  return (
    <div
      className={`${styles.container} ${shortenedUrl ? styles.expanded : ''}`}
    >
      <div className={styles.inputContainer}>
        <DefaultInput
          id="myInput"
          type="text"
          placeholder="Insira sua URL aqui"
          value={originalUrl}
          onChange={(e) => setOriginalUrl(e.target.value)}
        />
        <ShortenButton onClick={handleShortenUrl} />
      </div>
      {loading && <p>Gerando link...</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {shortenedUrl && <UrlOutput shortenedUrl={shortenedUrl} />}
    </div>
  )
}
