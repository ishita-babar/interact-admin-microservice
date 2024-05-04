import Toaster from '@/utils/toaster';
import Cookies from 'js-cookie';
import Image from 'next/image';
import { useRouter } from 'next/router';
import React from 'react';

const Header = () => {
  const router = useRouter();
  return (
    <div className="w-full bg-white p-4 flex justify-between items-center border-b-[1px] border-dashed border-primary_black">
      <div className="w-fit flex items-center">
        <div className="font-primary font-thin text-4xl pb-1">Logs</div>
      </div>
      <div
        onClick={() => {
          Cookies.remove('token');
          Cookies.remove('role');
          router.push('/login');
          Toaster.success('Logged Out');
        }}
        className="py-2 w-24 flex-center border-[1px] rounded-xl font-medium hover:scale-105 font-primary cursor-pointer transition-ease-300"
      >
        Logout
      </div>
    </div>
  );
};

export default Header;
