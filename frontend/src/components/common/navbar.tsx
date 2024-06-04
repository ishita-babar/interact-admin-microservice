import React from 'react';

const Navbar = () => {
  return (
    <div className="w-full h-navbar bg-navbar dark:bg-dark_navbar text-gray-500 dark:text-white border-gray-300 border-b-[1px] dark:border-0 backdrop-blur-sm fixed top-0 flex justify-between px-4 items-center z-20">
      <div className="font-primary font-thin text-4xl pb-1">Admin</div>
    </div>
  );
};

export default Navbar;
