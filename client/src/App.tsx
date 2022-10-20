import { useRoutes } from 'react-router-dom';
import reactLogo from './assets/react.svg';
import './App.css';
import { Home, Room } from './pages';

function App() {
  const routes = useRoutes([
    { path: '/', element: <Home /> },
    {
      path: '/room/',
      element: <Room />,
      children: [{ path: ':roomId', element: <Room /> }]
    }
  ]);
  return routes;
}

export default App;
