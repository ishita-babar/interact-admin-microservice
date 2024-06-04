import React, { useState } from 'react';
import SidebarItem from './sidebar_item';
import {
  ArrowLineLeft,
  DotsNine,
  EnvelopeSimpleOpen,
  Handshake,
  HouseLine,
  MegaphoneSimple,
  NoteBlank,
  RocketLaunch,
  Ticket,
  UserCircle,
  WarningCircle,
} from '@phosphor-icons/react';
import Cookies from 'js-cookie';
import Toaster from '@/utils/toaster';
import { useRouter } from 'next/router';

interface Props {
  index: number;
}

const Sidebar = ({ index }: Props) => {
  const [active, setActive] = useState(index);

  const router = useRouter();

  const handleLogout = () => {
    Cookies.remove('token');
    Cookies.remove('role');
    router.push('/login');
    Toaster.success('Logged Out');
  };
  return (
    <div className="w-sidebar h-base bg-sidebar border-gray-300 border-r-[1px] dark:border-0 dark:bg-dark_sidebar backdrop-blur-sm pt-[40px] fixed mt-navbar py-6 flex flex-col justify-between pl-[30px] transition-ease-out-500 max-lg:hidden">
      <div className="w-full flex flex-col gap-2">
        <SidebarItem
          index={0}
          title="Logs"
          url="/"
          icon={<HouseLine size={24} />}
          active={active}
          setActive={setActive}
        />
        <SidebarItem
          index={1}
          title="Reports"
          icon={<WarningCircle size={24} />}
          active={active}
          setActive={setActive}
        />
        <SidebarItem index={2} title="Feedbacks" icon={<Handshake size={24} />} active={active} setActive={setActive} />

        <div className="text-gray-500 font-medium p-[8.5px] pt-12">Flags</div>

        <SidebarItem index={3} title="Users" icon={<UserCircle size={24} />} active={active} setActive={setActive} />
        <SidebarItem index={4} title="Posts" icon={<NoteBlank size={24} />} active={active} setActive={setActive} />
        <SidebarItem
          index={5}
          title="Projects"
          icon={<RocketLaunch size={24} />}
          active={active}
          setActive={setActive}
        />
        <SidebarItem index={6} title="Events" icon={<Ticket size={24} />} active={active} setActive={setActive} />
        <SidebarItem
          index={7}
          title="Openings"
          icon={<EnvelopeSimpleOpen size={24} />}
          active={active}
          setActive={setActive}
        />
        <SidebarItem
          index={8}
          title="Announcements"
          icon={<MegaphoneSimple size={24} />}
          active={active}
          setActive={setActive}
        />
        <SidebarItem index={9} title="Polls" icon={<DotsNine size={24} />} active={active} setActive={setActive} />
      </div>

      <ArrowLineLeft
        onClick={handleLogout}
        className="cursor-pointer ml-2 mt-2 text-gray-500 dark:text-white transition-ease-500"
        size={24}
      />
    </div>
  );
};

export default Sidebar;
