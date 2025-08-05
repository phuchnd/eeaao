import React from 'react';
import styled from 'styled-components';
import './App.css';
import logo from './logo.png';
import { GoogleOAuthProvider, GoogleLogin } from '@react-oauth/google';

const App = () => {
  // Handle Google login success
  const handleGoogleSuccess = (credentialResponse) => {
    console.log('Google login success:', credentialResponse);
    // Here you would typically send the token to your backend for verification
  };

  // Handle Google login error
  const handleGoogleError = () => {
    console.error('Google login failed');
  };

  return (
    <AppContainer>
      <LoginContainer>
        <LogoContainer>
          <Logo src={logo} alt="Health Quest Logo" />
          <Title>Welcome to HealthQuest</Title>
          <IntroText>
            Your journey to better health begins here.
            Join fun, personalized challenges powered by your smart devices — walk, run, move — and earn unique NFT medals as proof of your achievements.
            Compete, improve, and unlock secret rewards through consistency and commitment.
            It's time to gamify your fitness, one quest at a time.
          </IntroText>
        </LogoContainer>

        <LoginForm>
          <SocialLoginSection>
            <GoogleOAuthProvider clientId={process.env.REACT_APP_GOOGLE_CLIENT_ID}>
              <GoogleLoginButton>
                <GoogleLogin
                  onSuccess={handleGoogleSuccess}
                  onError={handleGoogleError}
                  useOneTap
                  width="100%"
                />
              </GoogleLoginButton>
            </GoogleOAuthProvider>
          </SocialLoginSection>
        </LoginForm>
      </LoginContainer>
    </AppContainer>
  );
};

// Styled Components
const AppContainer = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 20px;
`;

const LoginContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  max-width: 450px;
  width: 100%;
  background-color: white;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  padding: 30px;

  @media (max-width: 768px) {
    padding: 20px;
    box-shadow: none;
    background-color: transparent;
  }
`;

const LogoContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 30px;
`;

const Logo = styled.img`
  width: 80px;
  height: 80px;
`;

const Title = styled.h1`
  font-size: 24px;
  color: #333;
  margin-top: 10px;
  font-weight: 500;
`;

const IntroText = styled.p`
  font-size: 16px;
  color: #555;
  text-align: center;
  margin-top: 15px;
  line-height: 1.5;
  max-width: 90%;
`;

const LoginForm = styled.form`
  width: 100%;
`;

const InputGroup = styled.div`
  margin-bottom: 20px;
`;

const Label = styled.label`
  display: block;
  margin-bottom: 8px;
  font-size: 14px;
  color: #555;
`;

const Input = styled.input`
  width: 100%;
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
  transition: border-color 0.3s;

  &:focus {
    border-color: #4285f4;
    outline: none;
  }

  &::placeholder {
    color: #aaa;
  }
`;

const ForgotPassword = styled.div`
  text-align: right;
  margin-bottom: 20px;
  font-size: 14px;
  color: #4285f4;
  cursor: pointer;

  &:hover {
    text-decoration: underline;
  }
`;

const LoginButton = styled.button`
  width: 100%;
  padding: 12px;
  background-color: #4285f4;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s;

  &:hover {
    background-color: #3367d6;
  }
`;


const SocialLoginSection = styled.div`
  display: flex;
  flex-direction: column;
  gap: 15px;
  width: 100%;
  margin-bottom: 20px;
`;

const GoogleLoginButton = styled.div`
  width: 100%;

  & > div {
    width: 100% !important;
  }
`;


export default App;
