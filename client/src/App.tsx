import { Container } from './components/Container'
import { Title } from './components/Title'
import { UrlShortenerForm } from './components/UrlShortenerForm'
import { LinkProvider } from './contexts/LinkContext/LinkContext'
import './styles/App.css'

export default function App() {
  return (
    <LinkProvider>
      <Container>
        <Title />

        <UrlShortenerForm />
      </Container>
    </LinkProvider>
  )
}

