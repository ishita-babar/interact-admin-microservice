import Navbar from '@/components/common/navbar';
import Head from 'next/head';
import React, { ReactNode } from 'react';

interface WrapperProps {
  title?: string;
  children: ReactNode;
}

const BaseWrapper: React.FC<WrapperProps> = ({ children, title = '' }) => {
  return (
    <>
      <Head>
        <title>{title} | Interact Admin</title>
      </Head>
      <Navbar />
      <div className="w-full flex">{children}</div>
    </>
  );
};

export default BaseWrapper;
