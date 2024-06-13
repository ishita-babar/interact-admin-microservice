import Link from 'next/link';
import React, { ReactNode } from 'react';

interface Props {
  title: string;
  icon: ReactNode;
  active: number;
  setActive: React.Dispatch<React.SetStateAction<number>>;
  index: number;
  url?: string;
  onClick?: () => void;
}

const SidebarItem = ({ title, icon, active, setActive, index, url = '', onClick }: Props) => {
  return (
    <Link
      href={`${url != '' ? url : '/' + title.toLowerCase()}`}
      onClick={() => {
        setActive(index);
        if (onClick) onClick();
      }}
      className={`w-[220px] h-10 p-[8.5px] rounded-lg ${
        active == index
          ? 'bg-primary_comp_hover text-primary_text dark:text-white dark:bg-[#0e0c2a59]'
          : 'hover:bg-primary_comp dark:hover:bg-[#0000002b] text-gray-500 dark:text-white'
      } relative font-primary font-medium items-center transition-ease-out-500`}
    >
      {icon}
      {<div className="absolute top-1/2 -translate-y-1/2 opacity-100 left-[64px] transition-ease-500">{title}</div>}
    </Link>
  );
};

export default SidebarItem;
