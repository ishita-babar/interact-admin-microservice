import configuredAxios from '@/config/axios';
import { User } from '@/types';
import Toaster from '@/utils/toaster';
import { ArrowRight, Eye, EyeClosed } from '@phosphor-icons/react';
import Cookies from 'js-cookie';
import { useRouter } from 'next/router';
import React, { useState } from 'react';
import Image from 'next/image';
import Head from 'next/head';

const SignUp = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [role, setRole] = useState('User');

  const [mutex, setMutex] = useState(false);

  const [showPassword, setShowPassword] = useState(false);

  const router = useRouter();

  const handleSubmit = async (el: React.FormEvent<HTMLFormElement>) => {
    el.preventDefault();
    if (mutex) return;
    setMutex(true);
    const formData = {
      username,
      password,
      confirmPassword: password,
      role,
    };
    const toaster = Toaster.startLoad('Creating an account');

    await configuredAxios
      .post(`${process.env.NEXT_PUBLIC_BACKEND_URL}/signup`, formData, {
        withCredentials: true,
      })
      .then(res => {
        if (res.status === 201) {
          Toaster.stopLoad(toaster, 'Account Created!', 1);
          const user: User = res.data.user;
          Cookies.set('token', res.data.token, {
            expires: Number(process.env.NEXT_PUBLIC_COOKIE_EXPIRATION_TIME),
          });
          Cookies.set('id', user.id, {
            expires: Number(process.env.NEXT_PUBLIC_COOKIE_EXPIRATION_TIME),
          });
          router.push('/');
        } else {
          if (res.data.message) Toaster.stopLoad(toaster, res.data.message, 0);
          else Toaster.stopLoad(toaster, 'Internal Server Error', 0);
        }
      })
      .catch(err => {
        if (err.response?.data?.message) Toaster.stopLoad(toaster, err.response.data.message, 0);
        else if (err.response?.data) Toaster.stopLoad(toaster, err.response.data, 0);
        else {
          Toaster.stopLoad(toaster, 'Internal Server Error', 0);
        }
        setMutex(false);
      });
  };

  return (
    <>
      <Head>
        <title>Sign Up</title>
      </Head>
      <div className="w-screen h-screen flex-center">
        <div className="w-full max-lg:w-full h-full font-primary py-8 px-8 flex flex-col justify-start items-center gap-32">
          <div className="w-full flex justify-start">
            <div className="w-fit flex items-center">
              <Image src={'/logo.png'} alt="" width={1000} height={1000} className="w-32 h-fit" />
              <div className="font-primary font-thin text-4xl pb-1">logs</div>
            </div>{' '}
          </div>
          <form onSubmit={handleSubmit} className="w-1/3 max-md:w-full flex flex-col items-center gap-6">
            <div className="flex flex-col gap-2 text-center">
              <div className="text-2xl font-semibold">Let&apos;s Get Started</div>
              <div className="text-gray-400">Time to log into interact logs</div>
            </div>

            <div className="w-full flex flex-col gap-4">
              <div className="flex flex-col gap-2">
                <div className="font-medium">Username</div>
                <input
                  name="username"
                  value={username}
                  onChange={el => setUsername(el.target.value)}
                  type="text"
                  className="w-full bg-white focus:outline-none border-2 p-2 rounded-xl text-gray-400"
                />
              </div>
              <div className="flex flex-col gap-2">
                <div className="font-medium">Password</div>
                <div className="w-full relative">
                  <input
                    name="password"
                    autoComplete="current-password"
                    value={password}
                    onChange={el => setPassword(el.target.value)}
                    type={showPassword ? 'text' : 'password'}
                    className="w-full bg-white p-2 rounded-xl focus:outline-none focus:bg-white border-2 text-gray-400 pr-10"
                  />
                  {showPassword ? (
                    <Eye
                      onClick={() => setShowPassword(false)}
                      className="absolute top-1/2 right-3 -translate-y-1/2 cursor-pointer"
                      size={20}
                      weight="regular"
                    />
                  ) : (
                    <EyeClosed
                      onClick={() => setShowPassword(true)}
                      className="absolute top-1/2 right-3 -translate-y-1/2 cursor-pointer"
                      size={20}
                      weight="regular"
                    />
                  )}
                </div>
                <div className="flex flex-col gap-2">
                  <div className="font-medium">Role</div>

                  <select
                    onChange={el => setRole(el.target.value)}
                    value={role}
                    className="px-1 py-3 w-full bg-white text-gray-400 border-2 rounded-xl focus:ring-0"
                  >
                    {['Member', 'Manager'].map((option, index) => (
                      <option
                        className="w-full bg-white focus:outline-none border-2 p-2 rounded-xl text-gray-400"
                        key={index}
                        value={option}
                      >
                        {option}
                      </option>
                    ))}
                  </select>
                </div>
              </div>
            </div>
            <div className="w-full p-1 flex flex-col gap-4 items-center">
              <button
                type="submit"
                className="w-full relative p-2 border-2 after:absolute after:-top-[3px] after:-left-[3px] after:-right-[3px] after:-bottom-[3.5px] after:-z-10 after:bg-[#395887] after:rounded-xl flex items-center cursor-pointer justify-center gap-2 bg-[#3D6DB3] hover:bg-[#345C98] active:bg-[#2D5185] border-[#d1d1d1a7] text-white py-2 rounded-xl font-semibold"
              >
                <div> Continue</div>
                <ArrowRight size={20} weight="regular" />
              </button>
              <div onClick={() => router.push('/login')} className="text-gray-400 text-sm cursor-pointer">
                Already have an Account? <span className="font-medium underline underline-offset-2">Login</span>
              </div>
            </div>
          </form>
        </div>
      </div>
    </>
  );
};

export default SignUp;
